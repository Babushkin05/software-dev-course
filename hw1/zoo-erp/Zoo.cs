using System.Collections;
using System;
using System.Runtime.InteropServices.Marshalling;
using System.Security.Cryptography.X509Certificates;

namespace zoo_erp
{
    public class Zoo
    {
        private int _zooKey;
        private IClinic _clinic;

        private int curId = 0;
        private Dictionary<int, ZooAnimal> animals = new Dictionary<int, ZooAnimal>();

        private Dictionary<int, ZooInventar> things = new Dictionary<int, ZooInventar>();

        public Zoo(IClinic clinic)
        {
            _clinic = clinic;
            Random rnd = new Random();
            _zooKey = rnd.Next();
        }

        public List<Herbo> getKindAnimals()
        {
            List<Herbo> kinds = new List<Herbo>();

            for (int i = 0; i < curId; ++i)
            {
                if (animals.ContainsKey(i) && animals[i].animal is Herbo)
                {
                    Herbo herboAnimal = animals[i].animal as Herbo;

                    if (herboAnimal.IsInterqactive())
                    {
                        kinds.Add(herboAnimal);
                    }
                }
            }

            return kinds;
        }

        public int CountFoodCompsuntion()
        {
            int res = 0;
            for(int i = 0; i < curId; ++i)
            {
                if(animals.ContainsKey(i)){
                    res += animals[i].animal.Food;
                }
            }

            return res;
        }

        public int AddAnimal(Animal animal)
        {
            if (_clinic.IsHealthy(animal))
            {
                ZooAnimal zooAnimal = new ZooAnimal(animal, curId, _zooKey);
                animals[curId] = zooAnimal;
                ++curId;
                return curId - 1;
            }
            else
                return -1;
        }

        public int AddThing(Thing thing)
        {
            ZooInventar zooInventar = new ZooInventar(thing, curId, _zooKey);
            things[curId] = zooInventar;
            ++curId;
            return curId - 1;
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