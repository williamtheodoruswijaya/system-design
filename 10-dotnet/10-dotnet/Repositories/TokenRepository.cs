using Microsoft.AspNetCore.Identity;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

namespace _10_dotnet.Repositories
{
    public class TokenRepository : ITokenRepository
    {
        private readonly IConfiguration configuration;
        public TokenRepository(IConfiguration configuration)
        {
            this.configuration = configuration;
        }
        public string CreateJWTToken(IdentityUser user, List<string> roles)
        {
            // step 1: buat object claim (ini bagian payload)
            var claims = new List<Claim>();

            // step 2: tambahin user info ke payload (user info yang mau dijadiin string)
            claims.Add(new Claim(ClaimTypes.Email, user.Email));
            foreach (var role in roles) claims.Add(new Claim(ClaimTypes.Role, role));

            // step 3: buat key-nya (ini buat bagian signature)
            var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(configuration["Jwt:Key"]));

            // step 4: buat credentials-nya (pake key yang tadi + algoritma hash-nya)
            var credentials = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

            // step 5: generate the token
            var token = new JwtSecurityToken(
                // ini dari appsettings.json buat bagian header (metadata)
                issuer: configuration["Jwt:Issuer"],
                audience: configuration["Jwt:Audience"],

                claims: claims,                             // masukin claims-nya
                expires: DateTime.Now.AddMinutes(30),       // masukin expired time-nya
                signingCredentials: credentials             // masukin credential (buat bagian signature)
            );

            // step 6: return the token as a string (using JwtSecurityTokenHandler)
            return new JwtSecurityTokenHandler().WriteToken(token);
        }

        /*
         * JWT Token Structure: Header.Payload.Signature
         * - Header: anggep aja metadata tentang token (algoritma hash-nya)
         * - Payload: claims (data tentang user yang di-encode jadi token)
         * - Signature: intinya buat validasi token-nya sesuai secret key
         */
    }
}
