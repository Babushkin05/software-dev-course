using System.Dynamic;

namespace zoo_erp
{
    /// <summary>
    /// Wrapper for items for zoo usage
    /// </summary>
    internal class ZooInventar : ZooOwnership
    {
        /// <summary>
        /// Wrapped item
        /// </summary>
        /// <value></value>
        public Thing thing { get; private set; }

        public ZooInventar(Thing thing_, int id, int zooKey) : base(zooKey, id)
        {
            thing = thing_;
        }

        public override string ToString()
        {
            return base.ToString() + thing.ToString();
        }
    }
}