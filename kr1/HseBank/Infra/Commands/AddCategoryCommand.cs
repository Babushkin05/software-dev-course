using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class AddCategoryCommand : ICommand
{
    private readonly ICategoryService _categoryService;
    private readonly Category _category;

    public AddCategoryCommand(ICategoryService categoryService, Category category)
    {
        _categoryService = categoryService;
        _category = category;
    }

    public object Execute()
    {
        _categoryService.AddCategory(_category);
        return new object();
    }
}