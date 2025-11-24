using _10_dotnet.Models.Domain;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Data
{
    public class DotNetDbContext: DbContext
    {
        // Pembuatan constructor ini basically intinya kek gini:
        /*
         * public class DotNetDbContext extends DbContext {
         *      public DotNetDbContext(DbContextOptions dbContextOptions) {
         *          super(dbContextOptions);
         *      }
         * }
         * 
         * jadi base(dbContextOptions itu manggil constructor parent class, which's constructor dari DbContext
         * ibaratnya manggil super() di Java.
         * 
         * Q: kenapa harus ada constructor ini?
         * A: Karena DbContext butuh options untuk konfigurasi koneksi database, provider database, dll. Yang bisa diatur di DbContextOptions dan di Inject di Main kita yaitu di Program.cs
         */
        public DotNetDbContext(DbContextOptions dbContextOptions): base(dbContextOptions)
        {
            
        }

        // DB Sets (Intinya semua entity/table yang ada di Database)
        public DbSet<Difficulty> Difficulties { get; set; }
        public DbSet<Region> Regions { get; set; }
        public DbSet<Walk> Walks { get; set; }
    }
}
