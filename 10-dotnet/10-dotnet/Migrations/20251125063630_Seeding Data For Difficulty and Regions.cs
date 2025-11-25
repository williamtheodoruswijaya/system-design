using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

#pragma warning disable CA1814 // Prefer jagged arrays over multidimensional

namespace _10_dotnet.Migrations
{
    /// <inheritdoc />
    public partial class SeedingDataForDifficultyandRegions : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.RenameColumn(
                name: "Id",
                table: "Walks",
                newName: "id");

            migrationBuilder.RenameColumn(
                name: "Id",
                table: "Difficulties",
                newName: "id");

            migrationBuilder.InsertData(
                table: "Difficulties",
                columns: new[] { "id", "Name" },
                values: new object[,]
                {
                    { new Guid("617d2eeb-3107-4cde-a56d-be0eb69c1ef1"), "Hard" },
                    { new Guid("64136753-eb5a-438a-ad3a-55ff87b1e9ba"), "Medium" },
                    { new Guid("d5044492-5690-4883-8e05-4fb0233e30de"), "Easy" }
                });

            migrationBuilder.InsertData(
                table: "Regions",
                columns: new[] { "id", "Code", "Name", "RegionImageUrl" },
                values: new object[,]
                {
                    { new Guid("6c09db5c-ec99-43e6-94db-3a8833fe4430"), "AKL", "Auckland", "https://wallpaperaccess.com/full/9352734.jpg" },
                    { new Guid("813f1efa-cbaf-4251-9809-20b695508ecd"), "WLG", "Wellington", "https://www.siliconera.com/wp-content/uploads/2022/06/bakemonogatarimanga.png" },
                    { new Guid("e6bf5231-0dce-42fe-be3d-2886e582e863"), "NSN", "Nelson", "https://pixelz.cc/wp-content/uploads/2019/12/lightning-returns-final-fantasy-xiii-uhd-4k-wallpaper.jpg" },
                    { new Guid("e924e018-7902-46c6-9a97-079b1305a580"), "STL", "Southland", null }
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Difficulties",
                keyColumn: "id",
                keyValue: new Guid("617d2eeb-3107-4cde-a56d-be0eb69c1ef1"));

            migrationBuilder.DeleteData(
                table: "Difficulties",
                keyColumn: "id",
                keyValue: new Guid("64136753-eb5a-438a-ad3a-55ff87b1e9ba"));

            migrationBuilder.DeleteData(
                table: "Difficulties",
                keyColumn: "id",
                keyValue: new Guid("d5044492-5690-4883-8e05-4fb0233e30de"));

            migrationBuilder.DeleteData(
                table: "Regions",
                keyColumn: "id",
                keyValue: new Guid("6c09db5c-ec99-43e6-94db-3a8833fe4430"));

            migrationBuilder.DeleteData(
                table: "Regions",
                keyColumn: "id",
                keyValue: new Guid("813f1efa-cbaf-4251-9809-20b695508ecd"));

            migrationBuilder.DeleteData(
                table: "Regions",
                keyColumn: "id",
                keyValue: new Guid("e6bf5231-0dce-42fe-be3d-2886e582e863"));

            migrationBuilder.DeleteData(
                table: "Regions",
                keyColumn: "id",
                keyValue: new Guid("e924e018-7902-46c6-9a97-079b1305a580"));

            migrationBuilder.RenameColumn(
                name: "id",
                table: "Walks",
                newName: "Id");

            migrationBuilder.RenameColumn(
                name: "id",
                table: "Difficulties",
                newName: "Id");
        }
    }
}
