using _10_dotnet.Models.Domain;
using _10_dotnet.Models.DTO;
using AutoMapper;

namespace _10_dotnet.Mappings
{
    public class AutoMapperProfiles:Profile
    {
        public AutoMapperProfiles()
        {
            CreateMap<Region, RegionDto>().ReverseMap();
            // .ForMember(x => x.Name, opt => opt.MapFrom(x => x.FullName)) ini kalau nama propertynya beda, tapi best-practices-nya sih sama persis ya.
            // .ReverseMap() kalau dari DTO <-- Domain, kalau gaada ya DTO --> Domain (tapi jadinya kalau ada .ReverseMap() bisa 2 arah)
            
            CreateMap<Region, AddRegionRequestDto>().ReverseMap();
            CreateMap<Region, UpdateRegionRequestDto>().ReverseMap();
        }
    }
}
