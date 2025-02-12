using System.Dynamic;
using System.Runtime.CompilerServices;

namespace zoo_erp
{
    /// <summary>
    /// Class for console interface of ERP
    /// </summary>
    public sealed partial class ConsoleApp
    {
        private Zoo _zoo;

        public ConsoleApp(Zoo zoo)
        {
            _zoo = zoo;
        }

        /// <summary>
        /// Function for run console app
        /// </summary>
        public void Run()
        {
            Console.WriteLine("Welcome to Moscow Zoo ERP!");
            Console.WriteLine("\nType something to continue :: ");
            Console.ReadKey();

            while (true)
            {
                int method = Slider("What do you want?",
                    new string[]{
                        "Add Animal",
                        "Count food consumption",
                        "Print kind animals",
                        "Exit console app"
                    }
                );
                switch(method){
                    case 0:
                        AddAnimal();
                        break;
                    case 1:
                        PrintConsumprion();
                        break;
                    case 2:
                        PrintKindAnimals();
                        break;
                    default:
                        return;
                }
            }
        }

        /// <summary>
        /// Method for adding animals to zoo
        /// </summary>
        private void AddAnimal()
        {
            int animalType = Slider("What animal you want to add?",
                new string[] {
                    "Monkey",
                    "Rabbit",
                    "Tiger",
                    "Wolf",
                    "Return"
                }
            );
            if(animalType == 4)
                return;
            
            Animal animal;

            string? name = "name";
            int food = 0;
            byte health = 0;
            byte kindness = 0;
            
            while(true){
                Console.WriteLine("\nGive the name :: ");
                name = Console.ReadLine();
                if(name is null)
                    continue;
                Console.WriteLine("\nType how much food this animal spend :: ");
                string? input = Console.ReadLine();
                if(!int.TryParse(input, out food))
                    continue;
                Console.WriteLine("\nType health of this animal :: ");
                input = Console.ReadLine();
                if(!byte.TryParse(input, out health))
                    continue;
                if(animalType < 2){
                    Console.WriteLine("\nType the kindness of this animal");
                    input = Console.ReadLine();
                    if(!byte.TryParse(input, out kindness))
                        continue;
                }
                break;
            }

            switch(animalType){
                case 0:
                    animal = new Monkey(name, food, health, kindness);
                    break;
                case 1:
                    animal = new Rabbit(name, food, health, kindness);
                    break;
                case 2:
                    animal = new Tiger(name, food, health);
                    break;
                case 3:
                    animal = new Wolf(name, food, health);
                    break;
                default:
                    throw new ArgumentException("Error in slader");
            }

            int id = _zoo.AddAnimal(animal);
            if(id != -1)
                Console.WriteLine($"\nAnimal with name {name} was succesfully added to zoo with id={id}");
            else
                Console.WriteLine($"\nUnable to add animal: {name} so ill((");
            
            Console.WriteLine("\nType something to continue :: ");
            Console.ReadKey();
        }

        /// <summary>
        /// Method to print summar consumption
        /// </summary>
        private void PrintConsumprion(){
            Console.WriteLine($"\nTotal food consumtion is {_zoo.CountFoodCompsuntion()} kg of food");

            Console.WriteLine("\nType something to continue :: ");
            Console.ReadKey();
        }

        /// <summary>
        /// Method for printing kind animals
        /// </summary>
        private void PrintKindAnimals(){
            List<Herbo> kinds = _zoo.getKindAnimals();
            Console.WriteLine($"\nThere is {kinds.Count} kind herbos :: ");
            foreach(Herbo herbo in kinds){
                Console.WriteLine(herbo);
            }

            Console.WriteLine("\nType something to continue :: ");
            Console.ReadKey();
        }
    }
}