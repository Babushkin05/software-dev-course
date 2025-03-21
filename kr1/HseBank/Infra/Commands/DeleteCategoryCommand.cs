using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class DeleteCategoryCommand : ICommand
{
    private readonly ICategoryService _categoryService;
    private readonly int _categoryId;

    public DeleteCategoryCommand(ICategoryService categoryService, int categoryId)
    {
        _categoryService = categoryService;
        _categoryId = categoryId;
    }

    public object Execute()
    {
        _categoryService.DeleteCategory(_categoryId);
        return new object();
    }
}