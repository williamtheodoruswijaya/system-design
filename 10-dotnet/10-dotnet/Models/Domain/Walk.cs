namespace _10_dotnet.Models.Domain
{
    public class Walk
    {
        public Guid Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public double LengthInKm { get; set; }
        public string? WalkImageUrl { get; set; }

        // Foreign Key
        public Guid DifficultyId { get; set; } // Basically Foreign Key dari tabel Difficulty
        public Guid RegionId { get; set; }

        // Navigation Property
        public Difficulty Difficulty { get; set; } // Basically ini buat nge-link ke entity Difficulty tapi harus kasih tau Foreign Key-nya di atas
        public Region Region { get; set; }
    }
}
