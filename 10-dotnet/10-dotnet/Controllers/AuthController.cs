using _10_dotnet.Models.DTO;
using _10_dotnet.Repositories;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity; // Identity itu literally services bawaan C# buat register dan login user
using Microsoft.AspNetCore.Mvc;

namespace _10_dotnet.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class AuthController : ControllerBase
    {
        private readonly UserManager<IdentityUser> userManager; // ini services dari Identity buat register dan login user
        private readonly ITokenRepository tokenRepository;
        public AuthController(UserManager<IdentityUser> userManager, ITokenRepository tokenRepository)
        {
            this.userManager = userManager;
            this.tokenRepository = tokenRepository;
        }

        [HttpPost]
        [Route("register")]
        public async Task<IActionResult> Register([FromBody] RegisterRequestDto registerRequestDto)
        {
            // step 1: create identity user object (NOTES: Kita basically define user object-nya (Domain) kayak gimana di sini)
            var identityUser = new IdentityUser
            {
                UserName = registerRequestDto.Username,
                Email = registerRequestDto.Username
            };

            // step 2: store the identity user object to database (pakai userManager)
            var identityResult = await userManager.CreateAsync(identityUser, registerRequestDto.Password);

            /*
             * Dari step 1 dan step 2 diatas, kita basically create user domain dengan attribut:
             * UserName
             * Email
             * Password
             */

            // step 3: add roles if the user is created successfully
            if (identityResult.Succeeded)
            {
                if (registerRequestDto.Roles != null && registerRequestDto.Roles.Any())
                {
                    identityResult = await userManager.AddToRolesAsync(identityUser, registerRequestDto.Roles); // Basically nambahin attribut baru yaitu Roles
                    if (identityResult.Succeeded)
                    {
                        return Ok("User registered successfully");
                    }
                }
            }

            return BadRequest("User registration failed");
        }

        [HttpPost]
        [Route("login")]
        public async Task<IActionResult> Login([FromBody] LoginRequestDto loginRequestDto)
        {
            // step 1: find the user by username
            var identityUser = await userManager.FindByEmailAsync(loginRequestDto.Username);
            if (identityUser == null)
            {
                return Unauthorized("Username doesn't exists!");
            }

            // step 2: check the password
            var isPasswordValid = await userManager.CheckPasswordAsync(identityUser, loginRequestDto.Password);
            if (!isPasswordValid)
            {
                return Unauthorized("Password is incorrect!");
            }

            // step 3: get roles for user (opsional, tapi ini butuh buat generate jwt token)
            var roles = await userManager.GetRolesAsync(identityUser);
            if (roles == null)
            {
                return Unauthorized("User has no roles assigned!");
            }

            // step 4: generate JWT token if the password is correct
            var jwtToken = tokenRepository.CreateJWTToken(identityUser, roles.ToList());

            // step 5: return the token to the client
            var response = new LoginResponseDto() // pake Dto biar bisa nambah atribut lain kalau mau
            {
                JwtToken = jwtToken
            };

            return Ok(response);
        }
    }
}
