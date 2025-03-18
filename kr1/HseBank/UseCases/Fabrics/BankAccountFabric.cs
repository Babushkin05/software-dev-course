using HseBank.Entities;

namespace HseBank.UseCases
{
    public static class BankAccountFabric
    {
        private static uint maxId = 0;

        public static BankAccount Create(string name){
            return new BankAccount(maxId++, name);
        }
    }
}