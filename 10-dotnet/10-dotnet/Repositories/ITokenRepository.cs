using Microsoft.AspNetCore.Identity;

namespace _10_dotnet.Repositories
{
    public interface ITokenRepository
    {
        string CreateJWTToken(IdentityUser user, List<string> roles);
    }
}
