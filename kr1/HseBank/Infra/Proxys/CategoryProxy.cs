using HseBank.UseCases;
using HseBank.Domain;

namespace HseBank.Infra;

public class CategoryServiceProxy : ICategoryService
{
    private readonly ICategoryService _realService;
    private readonly string _logFilePath;

    public CategoryServiceProxy(ICategoryService realService, string logFilePath)
    {
        _realService = realService;
        _logFilePath = logFilePath;
    }

    private void Log(string message)
    {
        string logMessage = $"{DateTime.Now}: {message}";
        File.AppendAllText(_logFilePath, logMessage + Environment.NewLine);
    }

    public void CreateCategory(Category category)
    {
        Log($"Creating category with ID {category.Id} and Name {category.Name}");
        try
        {
            _realService.CreateCategory(category);
            Log($"Category with ID {category.Id} created successfully.");
        }
        catch (Exception ex)
        {
            Log($"Error creating category: {ex.Message}");
            throw;
        }
    }

    public void AddCategory(Category category)
    {
        Log($"Adding category with ID {category.Id} and Name {category.Name}");
        try
        {
            _realService.AddCategory(category);
            Log($"Category with ID {category.Id} added successfully.");
        }
        catch (Exception ex)
        {
            Log($"Error adding category: {ex.Message}");
            throw;
        }
    }

    public void DeleteCategory(int id)
    {
        Log($"Deleting category with ID {id}");
        try
        {
            _realService.DeleteCategory(id);
            Log($"Category with ID {id} deleted successfully.");
        }
        catch (Exception ex)
        {
            Log($"Error deleting category: {ex.Message}");
            throw;
        }
    }

    public Category GetCategory(int id)
    {
        Log($"Getting category with ID {id}");
        try
        {
            var category = _realService.GetCategory(id);
            Log($"Retrieved category with ID {id}: {category?.Name ?? "Not Found"}");
            return category!;
        }
        catch (Exception ex)
        {
            Log($"Error getting category: {ex.Message}");
            throw;
        }
    }

    public IEnumerable<Category> GetAllCategories()
    {
        Log("Getting all categories");
        try
        {
            var categories = _realService.GetAllCategories();
            Log($"Retrieved {categories.Count()} categories.");
            return categories;
        }
        catch (Exception ex)
        {
            Log($"Error getting all categories: {ex.Message}");
            throw;
        }
    }
}