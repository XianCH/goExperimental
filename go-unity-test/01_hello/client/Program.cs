// See https://aka.ms/new-console-template for more information
using System;
using System.Net.Http;
using System.Threading.Tasks;

class Program
{
    static async Task Main(string[] args)
    {
        await SendHello();
    }

    static async Task SendHello()
    {
        HttpClient client = new HttpClient();
        HttpResponseMessage response = await client.PostAsync("http://localhost:12345/hello", null);
        
        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Response: " + responseContent);
        }
        else
        {
            Console.WriteLine("Request failed: " + response.StatusCode);
        }
    }
}
