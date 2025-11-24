using _10_dotnet.Data;
using _10_dotnet.Models.Domain;
using _10_dotnet.Models.DTO;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace _10_dotnet.Controllers
{
    [Route("api/[controller]")] // ini sama aja kayak "api/regions"
    [ApiController]
    public class RegionsController : ControllerBase
    {
        // Ini tempat buat Dependency Injection-nya nanti kalau kita butuh service, repository, dll. (termasuk DbContext)
        private readonly DotNetDbContext dbContext;
        public RegionsController(DotNetDbContext dbContext) // DbContext-nya di-inject via constructor
        {
            this.dbContext = dbContext;
        }

        [HttpGet]                       // Ini HTTP Method-nya ada [HttpPost], [HttpPut], [HttpDelete], dll. (Jadi, nanti cara akses-nya itu GET: https://localhost:portnumber/api/regions)
        public IActionResult GetAll()   // IActionResult itu return type khusus buat controller (Mirip kek kalau di NestJS/Golang Application, kita state response HTTP-nya dalam bentuk WebResponse<T>
        {
            // step 1: get data from database - Domain Models
            var regions = dbContext.Regions.ToList(); // Mengambil semua data Regions dari database (dalam bentuk Domain models)

            // step 2: map domain models to DTOs
            var regionsDto = new List<RegionDto>();
            foreach (var region in regions)
            {
                regionsDto.Add(new RegionDto()
                {
                    id = region.id,
                    Code = region.Code,
                    Name = region.Name,
                    RegionImageUrl = region.RegionImageUrl
                });
            }

            // step 3: return DTOs to client
            return Ok(regionsDto);                  // Ok() itu method dari ControllerBase yang nge-wrap response HTTP dengan status code 200
        }

        [HttpGet]
        [Route("{id:guid}")] // Ini contoh jenis Route dengan parameter.
                             // GET: https://localhost:portnumber/api/regions/{id}
                             // Sebenernya mau kode-nya [Route("{id}")] aja juga bisa dan aman aja
        public IActionResult GetById([FromRoute] Guid id)
        {
            // step 1: get data from database - Domain Models
            // Method 1: Pakai DBContext langsung
            // var region = dbContext.Regions.Find(id);

            // Method 2: Pakai LINQ
            var region = dbContext.Regions.FirstOrDefault(r => r.id == id);

            if (region == null)
            {
                return NotFound(); // NotFound() itu method dari ControllerBase yang nge-wrap response HTTP dengan status code 404
            }

            // step 2: map domain model to DTO
            var regionDto = new RegionDto()
            {
                id = region.id,
                Code = region.Code,
                Name = region.Name,
                RegionImageUrl = region.RegionImageUrl
            };

            // step 3: return DTO to client
            return Ok(regionDto);
        }

        [HttpPost]
        public IActionResult Create([FromBody] AddRegionRequestDto addRegionRequestDto) // buat ambil dari Body Request, kita pake [FromBody]
        {
            // step 1: map DTO to domain model
            var regionDomainModel = new Region()
            {
                Code = addRegionRequestDto.Code,
                Name = addRegionRequestDto.Name,
                RegionImageUrl = addRegionRequestDto.RegionImageUrl
            };

            // step 2: use domain model to create region (using DbContext or LINQ)
            dbContext.Regions.Add(regionDomainModel);

            // step 3: WAJIB! kalau di .NET, kita harus saveChanges() biar ke commit ke database
            dbContext.SaveChanges();

            // step 4: map created domain model to DTO
            var regionDto = new RegionDto()
            {
                id = regionDomainModel.id,
                Code = regionDomainModel.Code,
                Name = regionDomainModel.Name,
                RegionImageUrl = regionDomainModel.RegionImageUrl
            };

            // step 4: return created region to client
            return CreatedAtAction(nameof(GetById), new { id = regionDto.id }, regionDto);

            /*
             * Ini basically kalau Ok() itu status code 200,
             * Kalau CreatedAtAction() itu status code 201 (Created)
             * Terima 3 parameter:
             * - Nama action yang mau di-redirect (biasanya buat GET by id) -> ibaratnya nyari resource yang baru dibuat langsung di Database
             * - Object yang isinya parameter-parameter buat action yang dituju -> dalam kasus ini, GetById itu butuh parameter "id", jadi kita buat object baru yang isinya id dari region yang baru dibuat
             * - Object yang isinya data yang mau di-return ke client -> dalam kasus ini, kita return regionDto yang baru dibuat
             */
        }

        [HttpPut]
        [Route("{id:guid}")]
        public IActionResult Update([FromRoute] Guid id, [FromBody] UpdateRegionRequestDto regionRequestDto)
        {
            // step 1: cari region yang mau di-update
            var existingRegion = dbContext.Regions.FirstOrDefault(r => r.id == id);
            if (existingRegion == null)
            {
                return NotFound();
            }

            // step 2: update region yang ditemukan dengan data dari DTO
            existingRegion.Code = regionRequestDto.Code;
            existingRegion.Name = regionRequestDto.Name;
            existingRegion.RegionImageUrl = regionRequestDto.RegionImageUrl;

            // step 3: update region di database (langsung SaveChanges() aja karena entity-nya udah di-track oleh DbContext)
            dbContext.SaveChanges();

            // step 4: map updated domain model to DTO
            var regionDto = new RegionDto()
            {
                id = existingRegion.id,
                Code = existingRegion.Code,
                Name = existingRegion.Name,
                RegionImageUrl = existingRegion.RegionImageUrl
            };

            // step 5: return updated region to client
            return Ok();
        }
    }
}
