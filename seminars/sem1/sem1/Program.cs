using System;
namespace sem1
{
    public class Program
    {
        public static void Main()
        {
            FactoryAF factory = new FactoryAF(
                new Customer { FIO = "Gobets Ivan" },
                new Customer { FIO = "Klychkov Maxim" },
                new Customer { FIO = "Antohin Anon" },
                new Customer { FIO = "Vadim Zenin" },
                new Customer { FIO = "Dimas" }
            );

            for(int i = 0; i<3; ++i)
            {
                factory.Cars.Push(new Car());
            }

            Console.WriteLine("Before:");
            Console.WriteLine(string.Join(Environment.NewLine, factory.Cars));
            Console.WriteLine(string.Join(Environment.NewLine, factory.Customers));

            factory.SaleCar();

            Console.WriteLine("After:");
            Console.WriteLine(string.Join(Environment.NewLine, factory.Cars));
            Console.WriteLine(string.Join(Environment.NewLine, factory.Customers));
        }

    }
}