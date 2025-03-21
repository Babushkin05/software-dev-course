using System.Data.Common;

namespace HseBank.Domain
{
    public enum OperationType
    {
        Income,
        Consumption
    }
    public class Operation
    {
        public readonly UInt128 Id;

        public readonly OperationType Type;

        public readonly uint BankAccountId;

        public readonly uint Amount;

        public readonly DateTime Time;

        public readonly string? Description;

        public readonly uint CategoryId;

        internal Operation
        (
            UInt128 id_, OperationType type_,
            uint bankAccountId_, uint amount_,
            DateTime time_, string? description_,
            uint categoryId_
            )
        {
            Id = id_;
            Type = type_;
            BankAccountId = bankAccountId_;
            Amount = amount_;
            Time = time_;
            Description = description_;
            CategoryId = categoryId_;
        }

    }
}