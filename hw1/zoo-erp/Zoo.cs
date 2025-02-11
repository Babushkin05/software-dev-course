using System.Runtime.InteropServices.Marshalling;

namespace zoo_erp
{
    public class Zoo
    {
        private IClinic _clinic;

        private List<ZooAnimal> animals = new List<ZooAnimal>();

        private List<Thing> things = new List<Thing>();

        Zoo(IClinic clinic){
            _clinic = clinic;
        }

        public bool TryAddAnimal(Animal animal){
            if(_clinic.IsHealthy(animal)){
                ZooAnimal zooAnimal = new ZooAnimal(animal, animals.Count);           
                animals.Add(zooAnimal);
                return true;
            }
            else
                return false;
        }

        public void AddItem(Thing thing){
            thing
        }

    }
}