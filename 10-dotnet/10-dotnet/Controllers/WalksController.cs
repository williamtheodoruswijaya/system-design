using _10_dotnet.Models.Domain;
using _10_dotnet.Models.DTO;
using _10_dotnet.Repositories;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace _10_dotnet.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class WalksController : ControllerBase
    {
        private readonly IMapper mapper;
        private readonly IWalkRepository walkRepository;
        public WalksController(IMapper mapper, IWalkRepository walkRepository)
        {
            this.mapper = mapper;
            this.walkRepository = walkRepository;
        }

        [HttpPost]
        public async Task<IActionResult> Create([FromBody] AddWalkRequestDto addWalkRequestDto)
        {
            // step 1: map DTO to domain model
            var walkDomainModel = mapper.Map<Walk>(addWalkRequestDto);

            // step 2: save domain model to database
            var createdWalk = await walkRepository.CreateAsync(walkDomainModel);

            // step 3: map domain model to DTO
            var createdWalkDto = mapper.Map<WalkDto>(createdWalk);

            // step 4: return response
            return CreatedAtAction(nameof(GetById), new { id = createdWalkDto.id }, createdWalkDto);
        }

        [HttpGet]
        [Route("{id:guid}")]
        public async Task<IActionResult> GetById([FromRoute] Guid id)
        {
            // step 1: get domain model from database
            var walkDomainModel = await walkRepository.GetByIdAsync(id);

            // step 2: map domain model to DTO
            var walkDto = mapper.Map<WalkDto>(walkDomainModel);

            // step 3: return response
            return Ok(walkDto);
        }

        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            // step 1: get domain models from database
            var walkDomainModels = await walkRepository.GetAllAsync();

            // step 2: map domain models to DTOs
            var walkDtos = mapper.Map<List<WalkDto>>(walkDomainModels);

            // step 3: return response
            return Ok(walkDtos);
        }

        [HttpPut]
        [Route("{id:guid}")]
        public async Task<IActionResult> Update([FromRoute] Guid id, [FromBody] UpdateWalkRequestDto updateWalkRequestDto)
        {
            // step 1: map DTO to domain model
            var walkDomainModel = mapper.Map<Walk>(updateWalkRequestDto);

            // step 2: update domain model in database
            walkDomainModel = await walkRepository.UpdateAsync(id, walkDomainModel);
            if (walkDomainModel == null)
            {
                return NotFound();
            }

            // step 3: map updated domain model to DTO
            var updatedWalkDto = mapper.Map<WalkDto>(walkDomainModel);

            // step 4: return response
            return Ok(updatedWalkDto);
        }

        [HttpDelete]
        [Route("{id:guid}")]
        public async Task<IActionResult> Delete([FromRoute] Guid id)
        {
            // step 1: delete domain model from database
            var walkDomainModel = await walkRepository.DeleteAsync(id);
            if (walkDomainModel == null)
            {
                return NotFound();
            }
            
            // step 2: return response
            return Ok();
        }
    }
}
