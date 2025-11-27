using System.ComponentModel.DataAnnotations.Schema;

namespace _10_dotnet.Models.Domain
{
    public class Image
    {
        public Guid Id { get; set; }
        [NotMapped] // This property is not mapped to the database (jadi gaada di tabel)
        public IFormFile File { get; set; }
        public string FileName { get; set; }
        public string? FileDescription { get; set; }
        public string FileExtension { get; set; }
        public long FileSizeInBytes { get; set; }
        public string FilePath { get; set; }
    }
}
