using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public abstract class BankAccountImporter
{
    public List<BankAccount> ImportData(string filePath)
    {
        string fileContent = ReadFile(filePath);

        var accounts = ParseData(fileContent);

        return accounts;
    }

    protected virtual string ReadFile(string filePath)
    {
        return File.ReadAllText(filePath);
    }

    protected abstract List<BankAccount> ParseData(string fileContent);

}