using System.Collections;
using System;
using System.Runtime.InteropServices.Marshalling;

namespace zoo_erp
{
    public class Zoo
    {
        private int _zooKey;
        private IClinic _clinic;

        private Dictionary<int, ZooAnimal> animals = new Dictionary<int, ZooAnimal>();

        private Dictionary<int, ZooInventar> things = new Dictionary<int, ZooInventar>();

        Zoo(IClinic clinic)
        {
            _clinic = clinic;
            Random rnd = new Random();
            _zooKey = rnd.Next();
        }

        public bool TryAddAnimal(Animal animal)
        {
            if (_clinic.IsHealthy(animal))
            {
                ZooAnimal zooAnimal = new ZooAnimal(animal, animals.Count, _zooKey);
                animals[zooAnimal.Id] = zooAnimal;
                return true;
            }
            else
                return false;
        }

        public void AddItem(Thing thing)
        {
            ZooInventar zooInventar = new ZooInventar(thing, things.Count, _zooKey);
            things[zooInventar.Id] = zooInventar;
        }

        public Animal getAnimalById(int id)
        {
            return animals[id].animal;
        }

        public Thing GetThingById(int id)
        {
            return things[id].thing;
        }

        public Animal FreeAnimal(int id)
        {
            Animal animal = animals[id].animal;
            animals.Remove(id);
            return animal;
        }

        public Thing FreeThing(int id)
        {
            Thing thing = things[id].thing;
            things.Remove(id);
            return thing;
        }


        
    }
}