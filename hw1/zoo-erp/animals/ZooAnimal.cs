namespace zoo_erp
{
    internal class ZooAnimal : ZooOwnership
    {  
        public Animal animal { get; private set; }

        internal ZooAnimal(Animal animal_, int id, int zooKey) : base(zooKey, id)
        {
            animal = animal_;
        }

        public override string ToString()
        {
            return base.ToString() + animal.ToString();
        }
    }
}