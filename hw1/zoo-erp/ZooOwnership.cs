namespace zoo_erp
{
    /// <summary>
    /// Wrapper for ownership for zoo uses
    /// </summary>
    internal class ZooOwnership
    {
        /// <summary>
        /// Secret Key of owner
        /// </summary>
        private int _zooKey;

        /// <summary>
        /// Id in zoo ownership
        /// </summary>
        /// <value></value>
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