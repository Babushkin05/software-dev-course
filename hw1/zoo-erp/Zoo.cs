using System.Collections;
using System;
using System.Runtime.InteropServices.Marshalling;
using System.Security.Cryptography.X509Certificates;

namespace zoo_erp
{
    public class Zoo
    {
        /// <summary>
        /// Private password that uses for ruling ownership
        /// </summary>
        private int _zooKey;

        /// <summary>
        /// Clinic association
        /// </summary>
        private IClinic _clinic;

        /// <summary>
        /// Id for next ownership
        /// </summary>
        private int curId = 0;

        /// <summary>
        /// Database for animals
        /// </summary>
        /// <typeparam name="int">Index of animal</typeparam>
        /// <typeparam name="ZooAnimal">Animal</typeparam>
        /// <returns></returns>
        private Dictionary<int, ZooAnimal> animals = new Dictionary<int, ZooAnimal>();

        /// <summary>
        /// Database for itms
        /// </summary>
        /// <typeparam name="int">Index of item</typeparam>
        /// <typeparam name="ZooInventar">Item</typeparam>
        /// <returns></returns>
        private Dictionary<int, ZooInventar> things = new Dictionary<int, ZooInventar>();

        public Zoo(IClinic clinic)
        {
            _clinic = clinic;
            Random rnd = new Random();
            _zooKey = rnd.Next();
        }

        /// <summary>
        /// Return kind animals, that can be interactive
        /// </summary>
        /// <returns></returns>
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

        /// <summary>
        /// Return sum of all food consumption
        /// </summary>
        /// <returns></returns>
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

        /// <summary>
        /// Adding animal to database method
        /// </summary>
        /// <param name="animal">Animal to add</param>
        /// <returns></returns>
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

        /// <summary>
        /// Adding thing to database method
        /// </summary>
        /// <param name="thing"></param>
        /// <returns></returns>
        public int AddThing(Thing thing)
        {
            ZooInventar zooInventar = new ZooInventar(thing, curId, _zooKey);
            things[curId] = zooInventar;
            ++curId;
            return curId - 1;
        }

        /// <summary>
        /// Returns animal by id
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>

        public Animal getAnimalById(int id)
        {
            return animals[id].animal;
        }

        /// <summary>
        /// Returns item by id
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public Thing GetThingById(int id)
        {
            return things[id].thing;
        }

        /// <summary>
        /// Delete animal from database
        /// </summary>
        /// <param name="id"></param>
        /// <returns>animal</returns>
        public Animal FreeAnimal(int id)
        {
            Animal animal = animals[id].animal;
            animals.Remove(id);
            return animal;
        }

        /// <summary>
        /// Delete inventar from database
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public Thing FreeThing(int id)
        {
            Thing thing = things[id].thing;
            things.Remove(id);
            return thing;
        }
    }
}