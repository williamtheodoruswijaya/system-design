namespace _10_dotnet.Models.Domain
{
    public class Region
    {
        public Guid id { get; set; }
        public string Code { get; set; }
        public string Name { get; set; }
        public string? RegionImageUrl { get; set; } // string? basically artinya Nullable string type
    }
}
