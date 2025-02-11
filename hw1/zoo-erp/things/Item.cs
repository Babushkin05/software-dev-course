namespace zoo_erp
{
    public abstract class Thing : IInventory
    {
        public int ItemId { get; init; }

        public override string ToString()
        {
            return $"ItemId={ItemId}: ";
        }
    }
}