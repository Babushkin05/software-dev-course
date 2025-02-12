namespace zoo_erp
{
    /// <summary>
    /// Wrapper for animal for usage in zoo
    /// </summary>
    internal class ZooAnimal : ZooOwnership
    {
        /// <summary>
        /// Wrapped animal
        /// </summary>
        /// <value></value>
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