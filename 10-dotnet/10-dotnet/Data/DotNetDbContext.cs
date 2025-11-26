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
        public DotNetDbContext(DbContextOptions<DotNetDbContext> dbContextOptions): base(dbContextOptions)
        {
            
        }

        // DB Sets (Intinya semua entity/table yang ada di Database)
        public DbSet<Difficulty> Difficulties { get; set; }
        public DbSet<Region> Regions { get; set; }
        public DbSet<Walk> Walks { get; set; }

        // Data Seeding
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);

            // Seed data for Difficulty: arti-nya data awal yang bakal di-insert ke table Difficulty pas pertama kali database ini dibuat
            var difficulties = new List<Difficulty>()
            {
                new Difficulty()
                {
                    id = Guid.Parse("d5044492-5690-4883-8e05-4fb0233e30de"),
                    Name = "Easy"
                },
                new Difficulty()
                {
                    id = Guid.Parse("64136753-eb5a-438a-ad3a-55ff87b1e9ba"),
                    Name = "Medium"
                },
                new Difficulty()
                {
                    id = Guid.Parse("617d2eeb-3107-4cde-a56d-be0eb69c1ef1"),
                    Name = "Hard"
                },
            };

            // ini cara seed-nya:
            modelBuilder.Entity<Difficulty>().HasData(difficulties); // `modelBuilder.Entity<Model yang mau di seed ke database>().HasData(data yang mau di-seed);`

            // Seed data for Region
            var regions = new List<Region>()
            {
                new Region()
                {
                    id = Guid.Parse("6c09db5c-ec99-43e6-94db-3a8833fe4430"),
                    Name = "Auckland",
                    Code = "AKL",
                    RegionImageUrl = "https://wallpaperaccess.com/full/9352734.jpg"
                },
                new Region()
                {
                    id = Guid.Parse("813f1efa-cbaf-4251-9809-20b695508ecd"),
                    Name = "Wellington",
                    Code = "WLG",
                    RegionImageUrl = "https://www.siliconera.com/wp-content/uploads/2022/06/bakemonogatarimanga.png"
                },
                new Region()
                {
                    id = Guid.Parse("e6bf5231-0dce-42fe-be3d-2886e582e863"),
                    Name = "Nelson",
                    Code = "NSN",
                    RegionImageUrl = "https://pixelz.cc/wp-content/uploads/2019/12/lightning-returns-final-fantasy-xiii-uhd-4k-wallpaper.jpg"
                },
                new Region()
                {
                    id = Guid.Parse("e924e018-7902-46c6-9a97-079b1305a580"),
                    Name = "Southland",
                    Code = "STL",
                    RegionImageUrl = null
                }
            };
            modelBuilder.Entity<Region>().HasData(regions);
        }
    }
}
