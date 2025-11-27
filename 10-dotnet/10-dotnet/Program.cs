using _10_dotnet.Data;
using _10_dotnet.Mappings;
using _10_dotnet.Mappings;
using _10_dotnet.Repositories;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.IdentityModel.Tokens;
using System.Text;
using Microsoft.OpenApi.Models;
using Microsoft.Extensions.FileProviders;
using Serilog;
using _10_dotnet.Middlewares;

var builder = WebApplication.CreateBuilder(args);

// Add Logger Services
var logger = new LoggerConfiguration()
    .WriteTo.Console()
    .WriteTo.File("Logs/DotNet_Log.txt", rollingInterval : RollingInterval.Minute) // anggep aja buat nyatet log ke elastic search (ini ke file dulu)
    .MinimumLevel.Information()
    .CreateLogger();
builder.Logging.ClearProviders();
builder.Logging.AddSerilog(logger);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddHttpContextAccessor();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen(options => // Configure Swagger biar bisa Authorization pake JWT Token (COPAS AJA INI TEMPLATE BIAR SWAGGER BISA PAKE JWT AUTHORIZATION)
{
    options.SwaggerDoc("v1", new OpenApiInfo { Title = "10-dotnet", Version = "v1" });
    options.AddSecurityDefinition(JwtBearerDefaults.AuthenticationScheme, new OpenApiSecurityScheme
    {
        Name = "Authorization",
        In = ParameterLocation.Header,
        Type = SecuritySchemeType.ApiKey,
        Scheme = JwtBearerDefaults.AuthenticationScheme,
    });

    options.AddSecurityRequirement(new OpenApiSecurityRequirement
    {
        {
            new OpenApiSecurityScheme
            {
                Reference = new OpenApiReference
                {
                    Type = ReferenceType.SecurityScheme,
                    Id = JwtBearerDefaults.AuthenticationScheme
                },
                Scheme = "Oauth2",
                Name = JwtBearerDefaults.AuthenticationScheme,
                In = ParameterLocation.Header
            },
            new List<string>()
        }
    });
});


// Dependency Injection (Inject DBContext so DBContext is available throughout the application)
builder.Services.AddDbContext<DotNetDbContext>(options =>
{
    options.UseSqlServer(builder.Configuration.GetConnectionString("DotNetConnectionString"));
});
builder.Services.AddDbContext<DotNetAuthDbContext>(options =>
{
    options.UseSqlServer(builder.Configuration.GetConnectionString("DotNetAuthConnectionString"));
});

// Dependency Injection for Repositories (biar bisa diakses semua class/controller) <Interface, Implementation>
builder.Services.AddScoped<IRegionRepository, RegionRepository>();
builder.Services.AddScoped<IWalkRepository, WalkRepository>();
builder.Services.AddScoped<ITokenRepository, TokenRepository>();
builder.Services.AddScoped<IImageRepository, ImageRepository>();

// Dependency Injection buat AutoMapper Services
builder.Services.AddAutoMapper(config => {}, typeof(AutoMapperProfiles).Assembly);

// Dependency Injection for Identity Core Services (Nanti Login & Register bakal pakai services ini)
builder.Services.AddIdentityCore<IdentityUser>()
    .AddRoles<IdentityRole>()
    .AddTokenProvider<DataProtectorTokenProvider<IdentityUser>>("DotNet")
    .AddEntityFrameworkStores<DotNetAuthDbContext>()
    .AddDefaultTokenProviders();
builder.Services.Configure<IdentityOptions>(options =>
{
    options.Password.RequireDigit = false;
    options.Password.RequireLowercase = false;
    options.Password.RequireNonAlphanumeric = false;
    options.Password.RequireUppercase = false;
    options.Password.RequiredLength = 6;
    options.Password.RequiredUniqueChars = 1;
});

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

// Configure the HTTP request pipeline. (Basically kita masukin middleware custom dibawah semua ini)
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseMiddleware<ExceptionHandlerMiddleware>(); // Ini masukin custom middleware-nya biar global error handlingnya jalan

app.UseHttpsRedirection();

app.UseAuthentication(); // Ini biar JWT Authenticationnya jalan (sebelum authorization)

app.UseAuthorization();

app.UseStaticFiles(new StaticFileOptions
{
    FileProvider = new PhysicalFileProvider(Path.Combine(Directory.GetCurrentDirectory(), "Images")),
    RequestPath = "/Images"
}); // biar bisa akses file (.png, .css, dkk) di folder wwwroot

app.MapControllers();

app.Run();
