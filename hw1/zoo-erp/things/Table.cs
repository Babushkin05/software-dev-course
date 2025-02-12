namespace zoo_erp
{
    public class Table : Thing
    {
        public string? tableInfo;

        public Table(string? table)
        {
            tableInfo = table;
        }

        public override string ToString()
        {
            return base.ToString() + $" table {tableInfo}";
        }
    }
}