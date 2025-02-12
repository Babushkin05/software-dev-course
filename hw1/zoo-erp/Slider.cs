namespace zoo_erp
{
    public sealed partial class ConsoleApp{

        /// <summary>
        /// Slider for choosing something from a array.
        /// </summary>
        /// <param name="list">Array for variants.</param>
        /// <returns>Chosed index</returns>
        private static int Slider(string question, string[] list)
        {
            int isChoosen = -1;

            ConsoleKey pressed = ConsoleKey.UpArrow;
            do
            {
                Console.Clear();

                // Question.
                Console.WriteLine(question + '\n');

                // Process user activities.
                if (pressed == ConsoleKey.DownArrow)
                    isChoosen = Math.Min(isChoosen + 1, list.Length - 1);
                else if (pressed == ConsoleKey.UpArrow)
                    isChoosen = Math.Max(isChoosen - 1, 0);

                for (int i = 0; i < list.Length; i++)
                {
                    // Current chosed index has another theme.
                    if (i == isChoosen)
                    {
                        Console.ForegroundColor = ConsoleColor.White;
                        Console.WriteLine(list[i]);
                    }
                    else
                    {
                        Console.ForegroundColor = ConsoleColor.DarkGray;
                        Console.WriteLine(list[i]);
                    }
                }

                pressed = Console.ReadKey().Key;

                // Back to auto settings.
                Console.ResetColor();

            } while (pressed != ConsoleKey.Enter);

            return isChoosen;
        }
    }
}