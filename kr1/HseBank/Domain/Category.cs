namespace HseBank.Domain
{
    public enum CategoryType
    {
        Income,
        Consumption
    }

    public class Category
    {
        public readonly uint Id;

        public readonly CategoryType Type;

        public readonly string Name;

        internal Category(uint id_, CategoryType type_, string name_)
        {
            Id = id_;
            Type = type_;
            Name = name_;
        }

    }
}