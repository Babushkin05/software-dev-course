using HseBank.Domain;
namespace HseBank.UseCases;
public class CategoryService
{
    private readonly List<Category> _categories = new();

    public void CreateCategory(Category category)
    {
        if (category == null)
            throw new ArgumentNullException(nameof(category));

        if (_categories.Any(c => c.Id == category.Id))
            throw new InvalidOperationException("Category with the same ID already exists.");

        _categories.Add(category);
    }

    public void AddCategory(Category category)
    {
        if (category == null)
            throw new ArgumentNullException(nameof(category));

        _categories.Add(category);
    }

    public void DeleteCategory(int id)
    {
        var category = _categories.FirstOrDefault(c => c.Id == id);
        if (category == null)
            throw new InvalidOperationException("Category not found.");

        _categories.Remove(category);
    }

    public Category GetCategory(int id)
    {
        return _categories.FirstOrDefault(c => c.Id == id)!;
    }

    public IEnumerable<Category> GetAllCategories()
    {
        return _categories;
    }
}