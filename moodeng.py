import tkinter as tk

def create_window():
    # Créer la fenêtre principale
    new_window = tk.Tk()
    new_window.title("Alerte")
    label = tk.Label(new_window, text="Tu t'es fais moo denged", font=("Arial", 16))
    label.pack(pady=20)
    
    # Fermer la fenêtre et en ouvrir une autre lorsqu'on clique sur la croix
    def on_close():
        new_window.destroy()
        create_window()  # Appelle la fonction pour ouvrir une nouvelle fenêtre

    new_window.protocol("WM_DELETE_WINDOW", on_close)
    new_window.mainloop()

# Lancer la première fenêtre
create_window()
