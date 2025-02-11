namespace zoo_erp
{
    internal class ZooOwnership
    {
        private int _zooKey;

        private const int ITEM_HAVE_NO_OWNER = -1;


        public int Id { get; private set; }

        public ZooOwnership(int zooKey, int id)
        {
            _zooKey = zooKey;
            Id = id;
        }

        public override string ToString()
        {
            return $"ItemId={Id}: ";
        }
    }
}