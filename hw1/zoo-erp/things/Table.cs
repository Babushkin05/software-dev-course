namespace zoo_erp
{
    public class Table : Thing
    {
        /// <summary>
        /// Data about table
        /// </summary>
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