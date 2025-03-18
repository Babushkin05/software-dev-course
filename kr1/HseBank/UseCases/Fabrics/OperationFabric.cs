using HseBank.Entities;

namespace HseBank.UseCases
{
    public class OperationFabric
    {
        private static UInt128 maxId = 0;

        public Operation Create
        (
            OperationType type,
            uint bankAccountId,
            uint amount,
            string? description,
            uint categoryId
        )
        {
            return new Operation(
                maxId++, type, bankAccountId, amount, 
                DateTime.Now, description, categoryId
            );
        }
    }
}