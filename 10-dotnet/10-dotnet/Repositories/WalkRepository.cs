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
            var walks = await dbContext.Walks.ToListAsync();
            return walks;
        }

        public async Task<Walk?> GetByIdAsync(Guid id)
        {
            return await dbContext.Walks.FirstOrDefaultAsync(w => w.id == id);
        }
    }
}
