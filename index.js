var [wins, losses, abortions] = [0, 0, 0];
var gameInProcess, answer, maskedAnswer, wrongGuesses;
const masthead = document.querySelector("h1");
const commonWords = ["leur", "aurait", "à propos", "là", "penser", "lequel", "gens", "pourrait", "autre", "ces", "premier", "parce que", "chose", "ceux", "femme", "à travers", "enfant", "là", "après", "devrait", "monde", "école", "encore", "trois", "état", "jamais", "devenir", "entre", "vraiment", "quelque chose", "un autre", "famille", "partir", "pendant que", "étudiant", "formidable", "groupe", "commencer", "pays", "où", "problème", "chaque", "début", "pourrait", "à propos", "contre", "endroit", "encore", "entreprise", "où", "système", "droit", "programme", "question", "pendant", "gouvernement", "petit", "nombre", "toujours", "nuit", "point", "croire", "aujourd'hui", "apporter", "arriver", "sans", "avant", "grand", "million", "sous", "eau", "écrire", "mère", "national", "argent", "histoire", "jeune", "mois", "différent", "droit", "étude", "bien que", "entreprise", "problème", "noir", "petit", "maison", "après", "depuis", "fournir", "service", "autour", "ami", "important", "père", "jusqu'à", "puissance", "souvent", "politique", "parmi", "tenir", "cependant", "membre", "presque", "inclure", "continuer", "plus tard", "communauté", "blanc", "moins", "président", "apprendre", "changement", "minute", "plusieurs", "informations", "rien", "droit", "social", "comprendre", "si", "regarder", "ensemble", "suivre", "autour", "parent", "quelque chose", "créer", "public", "déjà", "parler", "autres", "niveau", "permettre", "bureau", "dépenser", "santé", "personne", "histoire", "parti", "dans", "résultat", "changer", "matin", "raison", "recherche", "tôt", "avant", "moment", "lui-même", "enseignant", "forcer", "offrir", "assez", "éducation", "à travers", "bien que", "se rappeler", "seconde", "peut-être", "vers", "politique", "tout", "processus", "musique", "y compris", "considérer", "apparaître", "en réalité", "probablement", "humain", "servir", "marché", "s'attendre à", "sens", "construire", "nation", "collège", "intérêt", "mort", "cours", "quelqu'un", "expérience", "derrière", "atteindre", "local", "rester", "effet", "suggérer", "classe", "contrôle", "augmenter", "peut-être", "peu", "champ", "ancien", "majeur", "parfois", "exiger", "le long de", "développement", "eux-mêmes", "rapport", "meilleur", "économique", "effort", "décider", "fort", "possible", "cœur", "leader", "lumière", "voix", "tout", "police", "enfin", "retour", "militaire", "prix", "rapport", "selon", "décision", "expliquer", "développer", "relation", "porter", "conduire", "fédéral", "pause", "mieux", "différence", "remercier", "recevoir", "valeur", "international", "bâtiment", "action", "modèle", "saison", "société", "parce que", "directeur", "tôt", "position", "joueur", "être d'accord", "surtout", "enregistrer", "papier", "spécial", "espace", "terrain", "soutien", "événement", "officiel", "dont", "matière", "tout le monde", "centre", "couple", "projet", "activité", "table", "tribunal", "produire", "enseigner", "situation", "industrie", "chiffre", "rue", "image", "lui-même", "téléphone", "soit", "couvrir", "assez", "image", "clair", "pratique", "pièce", "récent", "décrire", "produit", "médecin", "patient", "travailleur", "film", "certain", "nord", "personnel", "soutien", "simplement", "troisième", "technologie", "attraper", "ordinateur", "attention", "source", "presque", "organisation", "choisir", "cause", "point", "siècle", "preuve", "fenêtre", "difficile", "écouter", "culture", "milliard", "chance", "frère", "énergie", "période", "cours", "été", "réaliser", "centaine", "disponible", "plante", "probable", "opportunité", "court", "lettre", "condition", "choix", "endroit", "simple", "fille", "administration", "sud", "mari", "étage", "campagne", "matériel", "population", "économie", "médical", "hôpital", "église", "proche", "mille", "actuel", "futur", "mauvais", "impliquer", "défense", "quiconque", "augmentation", "sécurité", "moi-même", "certainement", "sport", "tableau", "sujet", "officier", "privé", "comportement", "performance", "combat", "lancer", "rapidement", "deuxième", "ordre", "auteur", "représenter", "mettre au point", "étranger", "sang", "agence"];
hideAll("#tally span");
document.querySelector("#new-game").addEventListener("click", newGame);

function newGame() {
if (gameInProcess == true) 
  aborted();
gameInProcess = true; 
masthead.innerText = "Hangman";
masthead.setAttribute("status", "normal");
answer = newRandomWord();
console.log("Hey you're cheating! " + 'Close the console! The answer is "' + answer + '"');
wrongGuesses = 0;
resetKeypad();
maskedAnswer = []; 
for (var i of answer)
  maskedAnswer.push("_");
updateDisplayWord(); 
hang();
}

function newRandomWord() {
return commonWords[Math.floor(Math.random() * commonWords.length)];
}

function verifyGuess() { 
var guessedLetter = this.innerText.toLowerCase();
if (answer.toLowerCase().includes(guessedLetter)) {
  for (var i in maskedAnswer) {
    if (answer[i] == guessedLetter)
      maskedAnswer[i] = answer[i];
  }
  updateDisplayWord();
  if (maskedAnswer.includes("_") == false)  
    escaped();
  this.classList.toggle("correct-letter", true);
  this.removeEventListener("click", verifyGuess);
} else {
  this.classList.toggle("incorrect-letter", true); 
  this.removeEventListener("click", verifyGuess);
  wrongGuesses++;
  
  hang();
}
}

function updateDisplayWord() {
var display = "";
for (var i of maskedAnswer)
  display += i + " ";
display.slice(0, -1);
document.querySelector("#guessing").textContent = display;
}

function aborted() { 
abortions++;
document.querySelector("#abortions").innerText = abortions;
unhideAll(".abortions");
}

function hang() {
switch (wrongGuesses) {
  case 0:
    hideAll("svg *");
    break;
  case 1:
    unhideAll(".gallows");
    break;
  case 2:
    unhide("#head");
    break;
  case 3:
    unhide("#body");
    break;
  case 4:
    unhide("#left-arm");
    break;
  case 5:
    unhide("#right-arm");
    break;
  case 6:
    unhide("#left-leg");
    break;
  case 7:
    unhide("#right-leg");
    hanged();
    break;
  default:
    newGame();
}
}


function hanged() { 
gameInProcess = false;
masthead.innerText = "Tu as perdu!";
masthead.setAttribute("status", "hanged");
losses++;
removeAllListeners();
unhideAll(".losses");
document.querySelector("#losses").innerText = losses;
var display = "";
for (var i of answer)
  display += i + " ";
display.slice(0, -1);
document.querySelector("#guessing").textContent = display;
}

function escaped() { 
gameInProcess = false;
masthead.innerText = "Tu as gagné!!";
masthead.setAttribute("status", "escaped");
wins++;
removeAllListeners();
unhideAll(".wins");
document.querySelector("#wins").innerText = wins;
}

function removeAllListeners() { 
for (let i of document.querySelectorAll("#keypad a")) {
  i.removeEventListener("click", verifyGuess);
  i.classList.toggle("finished", true);
}
}

function resetKeypad() {
for (var i of document.querySelectorAll("#keypad div")) 
  i.innerText = "";
populateRow(1, "ABCDEFGHI");
populateRow(2, "JKLMNOPQR");
populateRow(3, "STUVWXYZ");
}

function populateRow(rowNumber, rowLetters) { 
for (let i of rowLetters) {
  let key = document.createElement("a");
  key.id = i.toLowerCase();
  key.append(i);
  key.addEventListener("click", verifyGuess);
  document.querySelector("#keypad--row" + rowNumber).append(key);
}
}

function unhide(targetElement) {
document.querySelector(targetElement).classList.toggle("hidden", false);
}

function hideAll(targetElements) {
for (let i of document.querySelectorAll(targetElements))
  i.classList.toggle("hidden", true);
}

function unhideAll(targetElements) {
for (let i of document.querySelectorAll(targetElements))
  i.classList.toggle("hidden", false);
}
