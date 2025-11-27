using _10_dotnet.Models.Domain;

namespace _10_dotnet.Repositories
{
    public interface IImageRepository
    {
        Task<Image> UploadAsync (Image image);
    }
}
