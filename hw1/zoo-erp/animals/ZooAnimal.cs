namespace zoo_erp
{
    public class ZooAnimal
    {
        public Animal animal;

        public int animalId;

        public ZooAnimal(Animal animal_, int id)
        {
            animal = animal_;
            animalId = id;
        }

        public override string ToString()
        {
            return $"animalId={animalId}: " + animal.ToString();
        }
    }
}