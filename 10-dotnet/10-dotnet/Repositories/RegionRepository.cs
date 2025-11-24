using _10_dotnet.Data;
using _10_dotnet.Models.Domain;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Repositories
{
    public class RegionRepository: IRegionRepository
    {
        private readonly DotNetDbContext dbContext;
        public RegionRepository(DotNetDbContext dbContext)
        {
            this.dbContext = dbContext;
        }

        public async Task<Region> CreateAsync(Region region)
        {
            await dbContext.Regions.AddAsync(region);
            await dbContext.SaveChangesAsync();

            return region;
        }

        public async Task<Region?> Delete(Guid id)
        {
            var existingRegion = await dbContext.Regions.FirstOrDefaultAsync(r => r.id == id);
            if (existingRegion == null)
            {
                return null;
            }

            dbContext.Regions.Remove(existingRegion);
            await dbContext.SaveChangesAsync();

            return existingRegion;
        }

        public async Task<List<Region>> GetAllAsync()
        {
            return await dbContext.Regions.ToListAsync();
        }

        public async Task<Region?> GetByIdAsync(Guid id)
        {
            var region = await dbContext.Regions.FirstOrDefaultAsync(r => r.id == id);
            if (region == null)
            {
                return null;
            }
            return region;
        }

        public async Task<Region?> UpdateAsync(Guid id, Region region)
        {
            var existingRegion = await dbContext.Regions.FirstOrDefaultAsync(r => r.id == id);
            if (existingRegion == null)
            {
                return null;
            }

            existingRegion.Code = region.Code;
            existingRegion.Name = region.Name;
            existingRegion.RegionImageUrl = region.RegionImageUrl;

            await dbContext.SaveChangesAsync();
            return existingRegion;
        }
    }
}
