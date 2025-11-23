using _10_dotnet.Models.Domain;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Data
{
    public class DotNetDbContext: DbContext
    {
        public DotNetDbContext(DbContextOptions dbContextOptions): base(dbContextOptions)
        {
            
        }

        // DB Sets (Intinya semua entity/table yang ada di Database)
        public DbSet<Difficulty> Difficulties { get; set; }
        public DbSet<Region> Regions { get; set; }
        public DbSet<Walk> Walks { get; set; }
    }
}
