using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class CreateCategoryCommand : ICommand
{
    private readonly ICategoryService _categoryService;
    private readonly Category _category;

    public CreateCategoryCommand(ICategoryService categoryService, Category category)
    {
        _categoryService = categoryService;
        _category = category;
    }

    public object Execute()
    {
        _categoryService.CreateCategory(_category);
        return new object();
    }
}