using System.Dynamic;

namespace zoo_erp
{
    internal class ZooInventar : ZooOwnership
    {
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