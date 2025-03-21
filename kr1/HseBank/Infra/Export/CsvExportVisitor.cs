using HseBank.Domain;
using HseBank.UseCases;
using System.Text;

namespace HseBank.Infra;

public class CsvExportVisitor : IVisitor
{
    private readonly StringBuilder _csvContent = new();

    public void Visit(BankAccount account)
    {
        _csvContent.AppendLine($"{account.Id},{account.Name},{account.Balance}");
    }

    public void SaveToFile(string filePath)
    {
        File.WriteAllText(filePath, _csvContent.ToString());
    }
}