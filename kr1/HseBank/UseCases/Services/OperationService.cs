using HseBank.Domain;
namespace HseBank.UseCases;
public class OperationService
{
    private readonly List<Operation> _operations = new();

    public void CreateOperation(Operation operation)
    {
        if (operation == null)
            throw new ArgumentNullException(nameof(operation));

        if (_operations.Any(o => o.Id == operation.Id))
            throw new InvalidOperationException("Operation with the same ID already exists.");

        _operations.Add(operation);
    }

    public void DeleteOperation(UInt128 id)
    {
        var operation = _operations.FirstOrDefault(o => o.Id == id);
        if (operation == null)
            throw new InvalidOperationException("Operation not found.");

        _operations.Remove(operation);
    }

    public Operation GetOperation(UInt128 id)
    {
        return _operations.FirstOrDefault(o => o.Id == id)!;
    }

    public IEnumerable<Operation> GetAllOperations()
    {
        return _operations;
    }
}