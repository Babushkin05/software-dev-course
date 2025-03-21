namespace HseBank.Domain;

public interface IVisitor
{
    void Visit(BankAccount account);
}