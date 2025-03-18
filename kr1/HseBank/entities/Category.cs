namespace HseBank.entities
{
    public enum CategoryType
    {
        Income,
        Consumption
    }

    public class Category
    {
        public readonly uint id;

        public readonly CategoryType type;

        public readonly string name;

        internal Category(uint id_, CategoryType type_, string name_)
        {
            id = id_;
            type = type_;
            name = name_;
        }

    }
}