using System.ComponentModel.DataAnnotations;

namespace _10_dotnet.Models.DTO
{
    public class AddRegionRequestDto
    {
        [Required]
        [MinLength(3, ErrorMessage="Country code has to be a minimum of 3 Characters")] // MinLength(value, ErrorMessage="")
        [MaxLength(3, ErrorMessage="Country code has to be a maximum of 3 Characters")]
        public string Code { get; set; }

        [Required]
        [MaxLength(100, ErrorMessage ="Region name has to be a maximum of 100 Characters")]
        public string Name { get; set; }

        public string? RegionImageUrl { get; set; }
    }
}
