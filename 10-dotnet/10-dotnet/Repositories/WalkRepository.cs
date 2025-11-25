using _10_dotnet.Data;
using _10_dotnet.Models.Domain;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Repositories
{
    public class WalkRepository : IWalkRepository
    {
        private readonly DotNetDbContext dbContext;
        public WalkRepository(DotNetDbContext dbContext)
        {
            this.dbContext = dbContext;
        }

        public async Task<Walk> CreateAsync(Walk walk)
        {
            await dbContext.Walks.AddAsync(walk);
            await dbContext.SaveChangesAsync();
            return walk;
        }

        public async Task<List<Walk>> GetAllAsync()
        {
            return await dbContext.Walks
                .Include("Difficulty")
                .Include("Region")
                .ToListAsync(); 
            // .Include() basically performs a join operation between the related tables                                                                               // Kalau mau type-safe bisa pake x => x.Difficulty dan x => x.Region
        }

        public async Task<Walk?> GetByIdAsync(Guid id)
        {
            return await dbContext.Walks
                .Include("Difficulty")
                .Include("Region")
                .FirstOrDefaultAsync(w => w.id == id);
        }

        public async Task<Walk?> UpdateAsync(Guid id, Walk walk)
        {
            // step 1: get the existing walk from database
            var existingWalk = await dbContext.Walks.FirstOrDefaultAsync(w => w.id == id);
            if (existingWalk == null)
            {
                return null;
            }

            // step 2: update the properties
            existingWalk.Name = walk.Name;
            existingWalk.Description = walk.Description;
            existingWalk.LengthInKm = walk.LengthInKm;
            existingWalk.WalkImageUrl = walk.WalkImageUrl;
            existingWalk.DifficultyId = walk.DifficultyId;
            existingWalk.RegionId = walk.RegionId;

            // step 3: save the changes to database
            await dbContext.SaveChangesAsync();

            // step 4: return the updated walk
            return existingWalk;
        }
        public async Task<Walk?> DeleteAsync(Guid id)
        {
            var existingWalk = await dbContext.Walks.FirstOrDefaultAsync(w => w.id == id);
            if (existingWalk == null)
            {
                return null;
            }
            dbContext.Walks.Remove(existingWalk);
            await dbContext.SaveChangesAsync();

            return existingWalk;
        }
    }
}
