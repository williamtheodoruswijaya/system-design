using _10_dotnet.Data;
using _10_dotnet.Repositories;
using Microsoft.EntityFrameworkCore;

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
builder.Services.AddScoped<IRegionRepository, RegionRepository>(); // Ini buat dependency injection dari repository layer ke semua application

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}


app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
