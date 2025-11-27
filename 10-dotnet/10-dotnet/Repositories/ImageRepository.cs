using _10_dotnet.Data;
using _10_dotnet.Models.Domain;

namespace _10_dotnet.Repositories
{
    public class ImageRepository : IImageRepository
    {
        private readonly IWebHostEnvironment webHostEnvironment;
        private readonly IHttpContextAccessor httpContextAccessor;
        private readonly DotNetDbContext dbContext;

        public ImageRepository(
            IWebHostEnvironment webHostEnvironment, 
            IHttpContextAccessor httpContextAccessor, 
            DotNetDbContext dbContext)
        {
            this.webHostEnvironment = webHostEnvironment;
            this.httpContextAccessor = httpContextAccessor;
            this.dbContext = dbContext;
        }

        public async Task<Image> UploadAsync(Image image)
        {
            // step 1: create a local path
            var localFilePath = Path.Combine(
                webHostEnvironment.ContentRootPath, 
                "Images", 
                $"{image.FileName}{image.FileExtension}"
                );

            // step 2: save the file to the local path / storage (www.amazons3.com/your-bucket/images/image.png)
            using var stream = new FileStream(localFilePath, FileMode.Create);
            await image.File.CopyToAsync(stream);

            // step 3: set the file path so it can be accessed (https://localhost:1234/images/image.png)
            var urlFilePath = $"{httpContextAccessor.HttpContext.Request.Scheme}://" +
                              $"{httpContextAccessor.HttpContext.Request.Host}{httpContextAccessor.HttpContext.Request.PathBase}/" +
                              $"images/{image.FileName}{image.FileExtension}";

            // step 4: set the file path on the image object
            image.FilePath = urlFilePath;

            // step 5: save the image metadata to the database
            await dbContext.Images.AddAsync(image);
            await dbContext.SaveChangesAsync();

            // step 6: return the image object
            return image;
        }
    }
}
