using System.Data.Common;

namespace HseBank.entities
{
    public enum OperationType
    {
        Income,
        Consumption
    }
    public class Operation
    {
        public readonly UInt128 id;

        public readonly OperationType type;

        public readonly uint bankAccountId;

        public readonly uint amount;

        public readonly DateTime time;

        public readonly string? description;

        public readonly uint categoryId;

        internal Operation
        (
            UInt128 id_, OperationType type_,
            uint bankAccountId_, uint amount_,
            DateTime time_, string? description_,
            uint categoryId_
            )
        {
            id = id_;
            type = type_;
            bankAccountId = bankAccountId_;
            amount = amount_;
            time = time_;
            description = description_;
            categoryId = categoryId_;
        }

    }
}