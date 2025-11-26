using _10_dotnet.Data;
using _10_dotnet.Mappings;
using _10_dotnet.Mappings;
using _10_dotnet.Repositories;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.IdentityModel.Tokens;
using System.Text;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Dependency Injection (Inject DBContext so DBContext is available throughout the application)
builder.Services.AddDbContext<DotNetDbContext>(options =>
{
    options.UseSqlServer(builder.Configuration.GetConnectionString("DotNetConnectionString"));
});
builder.Services.AddDbContext<DotNetAuthDbContext>(options =>
{
    options.UseSqlServer(builder.Configuration.GetConnectionString("DotNetAuthConnectionString"));
});
builder.Services.AddScoped<IRegionRepository, RegionRepository>(); // Ini buat dependency injection dari repository layer ke semua application
builder.Services.AddScoped<IWalkRepository, WalkRepository>();
builder.Services.AddAutoMapper(config => {}, typeof(AutoMapperProfiles).Assembly);

// Dependency Injection for JWT Authentication Services (Copy-paste ini kalau mau pake JWT Authentication)
builder.Services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
    .AddJwtBearer(options =>
    options.TokenValidationParameters = new TokenValidationParameters
    {
        ValidateIssuer = true,
        ValidateAudience = true,
        ValidateLifetime = true,
        ValidateIssuerSigningKey = true,
        ValidIssuer = builder.Configuration["Jwt:Issuer"],
        ValidAudience = builder.Configuration["Jwt:Audience"],
        IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(builder.Configuration["Jwt:Key"]))
    });

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthentication(); // Ini biar JWT Authenticationnya jalan (sebelum authorization)

app.UseAuthorization();

app.MapControllers();

app.Run();
