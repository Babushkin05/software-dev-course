using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class GetCategoryCommand : ICommand
{
    private readonly ICategoryService _categoryService;
    private readonly int _categoryId;

    public GetCategoryCommand(ICategoryService categoryService, int categoryId)
    {
        _categoryService = categoryService;
        _categoryId = categoryId;
    }

    public object Execute()
    {
        return _categoryService.GetCategory(_categoryId);
    }
}