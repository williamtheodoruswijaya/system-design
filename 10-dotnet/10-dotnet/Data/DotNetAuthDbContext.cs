using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Data
{
    public class DotNetAuthDbContext: IdentityDbContext
    {
        public DotNetAuthDbContext(DbContextOptions<DotNetAuthDbContext> options) : base(options) // <- basically inject DbContext with our options configured in Program.cs + DbContextOptions<DotNetAuthDbContext> harus kek gini buat setiap DbContext kalau kita pake lebih dari 1 DbContext.
        {

        }

        // Data Seeding (Roles)
        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);

            var readerRoleId = "db27e761-cafb-410d-8cad-41d6d0c238c2";
            var writerRoleId = "d1722179-32d0-431c-9127-4b74f100b1f7";
            var roles = new List<IdentityRole>
            {
                new IdentityRole
                {
                    Id = readerRoleId,
                    ConcurrencyStamp = readerRoleId,
                    Name = "Reader",
                    NormalizedName = "Reader".ToUpper()
                },
                new IdentityRole
                {
                    Id = writerRoleId,
                    ConcurrencyStamp = writerRoleId,
                    Name = "Writer",
                    NormalizedName = "Writer".ToUpper()
                }
            };

            // Execute the seeding to the Roles table in the database
            builder.Entity<IdentityRole>().HasData(roles);

            /*
             * Cara Execute-nya:
             * 1. Add-Migration "Nama Message Migration" -Context "Nama DbContext yang mau di migrasi"
             * 2. Update-Database -Context "Nama DbContext yang mau di migrasi"
             * 
             * 
             * Notes:
             * Ini kita pakai IdentityRole buat seeding. Nah ini tuh apa?
             * Basically, IdentityRole itu adalah class yang sudah disediakan oleh ASP.NET Core Identity buat meng-handle role management.
             * Caranya, dia buat table-table di Database baru (Sesuai sama DbContext yang kita pake)
             * Dan tabel-tabel yang dia buat tuh, ada:
             *  - AspNetRoles      : Menyimpan daftar nama role (contoh: "Admin", "User", "Mitra").
             *  - AspNetUserRoles  : Tabel penghubung (Junction Table) many-to-many
             */
        }
    }
}
