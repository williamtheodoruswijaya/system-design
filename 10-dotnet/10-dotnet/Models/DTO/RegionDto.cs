namespace _10_dotnet.Models.DTO
{
    public class RegionDto // Basically subset dari Domain models (kasihin atribut yang mau di-expose ke client aja)
    {
        public Guid id { get; set; }
        public string Code { get; set; }
        public string Name { get; set; }
        public string? RegionImageUrl { get; set; }
    }
}
