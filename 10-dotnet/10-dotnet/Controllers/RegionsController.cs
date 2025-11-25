using _10_dotnet.CustomActionFilters;
using _10_dotnet.Data;
using _10_dotnet.Models.Domain;
using _10_dotnet.Models.DTO;
using _10_dotnet.Repositories;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace _10_dotnet.Controllers
{
    [Route("api/[controller]")] // ini sama aja kayak "api/regions"
    [ApiController]
    public class RegionsController : ControllerBase
    {
        // Ini tempat buat Dependency Injection-nya nanti kalau kita butuh service, repository, dll. (termasuk DbContext)
        private readonly IRegionRepository regionRepository;
        private readonly IMapper mapper;
        public RegionsController(IRegionRepository regionRepository, IMapper mapper) // DbContext-nya di-inject via constructor
        {
            this.regionRepository = regionRepository;
            this.mapper = mapper;
        }

        [HttpGet]                       // Ini HTTP Method-nya ada [HttpPost], [HttpPut], [HttpDelete], dll. (Jadi, nanti cara akses-nya itu GET: https://localhost:portnumber/api/regions)
        public async Task<IActionResult> GetAll()   // IActionResult itu return type khusus buat controller (Mirip kek kalau di NestJS/Golang Application, kita state response HTTP-nya dalam bentuk WebResponse<T>
                                                    // Kalau async, return type-nya harus di dalam Task<T>
                                                    // Function-function DbContext juga tinggal kasih keyword "await" dan cari yang versi Async-nya.
        {
            // step 1: get data from database - Domain Models
            var regions = await regionRepository.GetAllAsync();

            // step 2: map domain models to DTOs
            //var regionsDto = new List<RegionDto>();
            //foreach (var region in regions)
            //{
            //    regionsDto.Add(new RegionDto()
            //    {
            //        id = region.id,
            //        Code = region.Code,
            //        Name = region.Name,
            //        RegionImageUrl = region.RegionImageUrl
            //    });
            //}
            var regionsDto = mapper.Map<List<RegionDto>>(regions);

            // step 3: return DTOs to client
            return Ok(regionsDto);                  // Ok() itu method dari ControllerBase yang nge-wrap response HTTP dengan status code 200
        }

        [HttpGet]
        [Route("{id:guid}")] // Ini contoh jenis Route dengan parameter.
                             // GET: https://localhost:portnumber/api/regions/{id}
                             // Sebenernya mau kode-nya [Route("{id}")] aja juga bisa dan aman aja
        public async Task<IActionResult> GetById([FromRoute] Guid id)
        {
            // step 1: get data from database - Domain Models
            // Method 1: Pakai DBContext langsung
            // var region = dbContext.Regions.Find(id);

            // Method 2: Pakai LINQ
            // var region = await dbContext.Regions.FirstOrDefaultAsync(r => r.id == id);

            var region = await regionRepository.GetByIdAsync(id);
            if (region == null)
            {
                return NotFound(); // NotFound() itu method dari ControllerBase yang nge-wrap response HTTP dengan status code 404
            }

            // step 2: map domain model to DTO
            //var regionDto = new RegionDto()
            //{
            //    id = region.id,
            //    Code = region.Code,
            //    Name = region.Name,
            //    RegionImageUrl = region.RegionImageUrl
            //};
            var regionDto = mapper.Map<RegionDto>(region); // Intinya, Map<Destination>(Source)

            // step 3: return DTO to client
            return Ok(regionDto);
        }

        [HttpPost]
        [ValidateModel]
        public async Task<IActionResult> Create([FromBody] AddRegionRequestDto addRegionRequestDto) // buat ambil dari Body Request, kita pake [FromBody]
        {
            // step 1: map DTO to domain model
            //var regionDomainModel = new Region()
            //{
            //    Code = addRegionRequestDto.Code,
            //    Name = addRegionRequestDto.Name,
            //    RegionImageUrl = addRegionRequestDto.RegionImageUrl
            //};
            var regionDomainModel = mapper.Map<Region>(addRegionRequestDto);

            // step 2: use domain model to create region (using DbContext or LINQ)
            // step 3: WAJIB! kalau di .NET, kita harus saveChanges() biar ke commit ke database
            regionDomainModel = await regionRepository.CreateAsync(regionDomainModel);

            // step 4: map created domain model to DTO
            //var regionDto = new RegionDto()
            //{
            //    id = regionDomainModel.id,
            //    Code = regionDomainModel.Code,
            //    Name = regionDomainModel.Name,
            //    RegionImageUrl = regionDomainModel.RegionImageUrl
            //};
            var regionDto = mapper.Map<RegionDto>(regionDomainModel);

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
        [ValidateModel]
        public async Task<IActionResult> Update([FromRoute] Guid id, [FromBody] UpdateRegionRequestDto regionRequestDto)
        {
            // step 0: map DTO to domain model
            //var regionDomainModel = new Region()
            //{
            //    Code = regionRequestDto.Code,
            //    Name = regionRequestDto.Name,
            //    RegionImageUrl = regionRequestDto.RegionImageUrl
            //};
            var regionDomainModel = mapper.Map<Region>(regionRequestDto);

            // step 1: cari region yang mau di-update
            // step 2: update region yang ditemukan dengan data dari DTO
            // step 3: update region di database (langsung SaveChanges() aja karena entity-nya udah di-track oleh DbContext)
            var updatedRegion = await regionRepository.UpdateAsync(id, regionDomainModel);

            // step 4: map updated domain model to DTO
            var regionDto = mapper.Map<RegionDto>(updatedRegion);

            // step 5: return updated region to client
            return Ok(regionDto);
        }

        [HttpDelete]
        [Route("{id:guid}")]
        public async Task<IActionResult> Delete([FromRoute] Guid id)
        {
            // step 1: find existing region that want to delete
            // step 2: delete region
            // khusus delete, DbContext ga punya async method-nya jadi gaush dijadiin await.
            // step 3: save changes to database
            var deletedRegion = await regionRepository.Delete(id);
            if (deletedRegion == null)
            {
                return NotFound();
            }

            // step 4: return success message to client
            return Ok();
        }
    }
}
