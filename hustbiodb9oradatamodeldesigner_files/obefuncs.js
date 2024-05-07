/*
 * Oracle By Example (OBE): objfuncs.js JavaScript Document
 * Author: Glenn Stokol, Sergiy Pecherskyy, Marcie Young
 * History
 * =======
 * 23-Mar-09	GS	Created initial file from main OBE content
 * 02-Apr-09	GS	Added obe_link class for print preview and showall hideall images. Fixed toggleAllImages()
 *                  behavior. Fixed imgborder issue, and navigation bar title color.
 * 05-Apr-09	GS	Fixed Print Preview link to display in new browser window.
 */
var obeImages = new Array();
var eyeglass = new Image();
var topics = new Array();
var printPreview = false;
var showImages = false;
var isNav = (navigator.appName == "Netscape");

function obeInitialize()
{
  checkParameters();
  preLoadImages();
  createTopics(topics);  
  createNavbar();
  
  if (!printPreview) { 
     createTOC(topics);
	 hideMainTopics();
  }
  else {
     createMainTopics();
     createSubTopicSections();
  }
  if (!showImages) hideAllImages();
}

function checkParameters()
{
  if (location.search.length > 0) {
    var params = location.search.substr(1).split("&");
    for (var i=0; i<params.length; i++)
    {
	  var vals = params[i].split("=");
	  switch (vals[0]) {
	    case "print": printPreview = (vals[1] == "preview"); break;
		case "imgs":showImages =(vals[i] == "visible"); break;
	    default: break; //ignore all other parameters
	  }
    }
  }
}

function createNavbar()
{
  var navbarDiv = document.getElementById("obe_nav");
  var navbarHTML = "";
  
  if (!printPreview)
  {
    var topTitle = document.getElementById("top");
    navbarHTML += createPrintLink();
    navbarHTML += "<h1 class=\"obe_section\">" +
	              "<a href=\"#top\">" + (topTitle ? topTitle.innerHTML : "OBE Title Error") + "</a></h1>";
	//navbarHTML += "<h2 class=\"navheader\">Topics</h2>";
	navbarHTML += "<p>This tutorial contains the following sections:</p>";
	navbarHTML += "<div class=\"navlink\" id=\"obe_toc\"></div>";
	navbarHTML += "<h4 class=\"obe_section\">Viewing Screenshots</h4>"; // class="navheader subnavlinkpipe"
	navbarHTML += createHideShowAll();  //class="subnavlinkpipe"
	navbarDiv.className = "obe_navbar subnavlinkpipe"; // "
	document.getElementById("obe_content").className = "obe_content";
  }
  else
  {
    navbarHTML += createPrintLink();
	navbarDiv.className = "obe_navbar_printview no_print";
	document.getElementById("obe_content").className = "obe_content_printview";
  }
  navbarDiv.innerHTML = navbarHTML;
}

function createHideShowAll()
{
 	var printLink = "";
		printLink += "<p><img class=\"obe_link\" id=\"toggleallimg\" src=\"library/showall.gif\" align=\"absmiddle\" onClick=\"javascript:toggleAllImages(this);\" />"
		printLink += "<span id=\"togglealltext\" class=\"obe_nav_title\">Click icon to show all screenshots</span></p>"
		printLink += "<p><b>Note: </b>Alternatively, you can click an individual icon (or image) associated with each step to view (or hide) the screenshot associated with that step.</p>";
	return printLink;
}

function setPrintPreview(obj)
{
  var newLoc = "?print=preview";
  var toggleImg = document.getElementById("toggleallimg");
  // if the src is currently hideall, then the images are visible, and append this as a parameter
  if (toggleImg && toggleImg.src.indexOf("hideall") > -1) { 
    newLoc += "&imgs=visible";
  } 
  obj.href = newLoc;
  return true;
}

function createPrintLink()
{
 	var printLink = "";
	if (printPreview)
	{
  	    var imgSrc = (showImages ? "hideall" : "showall");
	    var imgTxt = "Click icon to " + (showImages ? "hide" : "show") + " all screenshots";
	    
		printLink += "<table width=\"100%\"><tr><td align=\"left\">"
		printLink += "<img id=\"toggleallimg\" class=\"obe_link\" src=\"library/" + imgSrc + ".gif\" align=\"absmiddle\" onClick=\"javascript:toggleAllImages(this);\" />"
		printLink += "<span id=\"togglealltext\" class=\"obe_nav_title\">" + imgTxt + "</span>";
		printLink += "</td><td align=\"right\">";
		printLink += "&nbsp;<a class=\"obe_link\" onClick=\"javascript:window.print();return false;\">";
		printLink += "<img src=\"library/print_17x15.gif\" width=\"17\" height=\"15\"  align=\"absmiddle\" /> Print</a>";
		printLink += "</td></tr></table>";		
	}
	else
	{
		printLink += "<a target=\"_blank\" class=\"obe_link\" href=\"\" onClick=\"javascript:setPrintPreview(this);\">";		
		printLink += "<img class=\"obe_link\" src=\"library/print_17x15.gif\" width=\"17\" height=\"15\" /> Print Preview</a>";
	}
	return printLink;
}

function Topic(n,t) {
	this.name = n;
	this.text = t;
	this.subtopics = null;
}

function printMainTopics(tarray)
{
	var toc = printTopics(tarray, printPreview);
	return toc;
}

function printSubTopics(tarray, tname)
{
	var toc = "";
	for (var i=0; i<tarray.length;i++) 
	{
	  if (tarray[i].name == tname) {
	     toc += printTopics(tarray[i].subtopics, false);
		 break;
	  }
	}
	return toc;
}

function printTopics(tarray, inc_subtopics)
{
	var toc = "";
	toc += "<table>";
	for (var i=0; i<tarray.length;i++) 
	{
		toc += "<tr>";
		if (inc_subtopics && tarray[i].subtopics) {
		  toc += "<td align='center' valign='middle'><img src='library/expand.gif' alt='expand'"
		  toc += " onclick=\"javascript:toggleSubtopics(this);toggleVisible('sub_" + tarray[i].name + "');\" /></td>"; // bullet
		}
		else {
//		  toc += "<td><img src=\"library/bullet.gif\" alt=\"bullet\"/></td>"; // bullet
		  toc += "<td></td>"; // MEY 3/17/09 no bullet
		}
		toc += "<td>";
		toc += tarray[i].text.link("#"+tarray[i].name);
		toc += "</td>";
		toc += "</tr>";
		if (inc_subtopics && tarray[i].subtopics) {
    		toc += "<tr id='sub_" + tarray[i].name + "' style=\"display:none;\">"; // add new row for subtopics
			toc += "<td></td>"; // empty 1st col for subtopics 
			toc += "<td>" 
			toc += printTopics(tarray[i].subtopics);
		    toc += "</td>";
		    toc += "</tr>";
		}
	}
	toc += "</table>"
	return toc;
}

function hideMainTopics()
{
  var mainTopicDiv = document.getElementById("obe_maintopics");
  if (mainTopicDiv) {
    mainTopicDiv.style.display = "none";
  }
}

function createMainTopics()
{
  var mainTopicDiv = document.getElementById("obe_maintopics");
  // var divHTML = "<h2 class=\"navheader\">Topics</h2><p>This tutorial contains the following sections:</p>";
  var divHTML = "<p>This tutorial contains the following sections:</p>";

  if (mainTopicDiv) {
     divHTML += "<div id=\"obe_toc\" class=\"obe_toc_printview\">";
     divHTML += printMainTopics(topics);
	 divHTML += "</div>";
     mainTopicDiv.innerHTML = divHTML;
     mainTopicDiv.style.display = "block";
  }
} 

function insertContentBefore(node, divname)
{
  var myElement = document.createElement("DIV");
  var newContent = printSubTopics(topics, divname);
  myElement.id = divname + "_div";
  myElement.innerHTML =newContent;
  if (node.parentNode) {
     node.parentNode.insertBefore(myElement, node);
  }
}

function createSubTopicSections()
{
  for (var i=0; i < topics.length; i++)
  {
    if (topics[i].subtopics && topics[i].subtopics.length > 0) {
	   var theId = topics[i].subtopics[0].name;
	   var firstSubtopicHead = document.getElementById(theId);  // This should be the first subtopic anchor
	   if (firstSubtopicHead.parentNode)  { // check that the first subtopic parent node exists
	     insertContentBefore(firstSubtopicHead.parentNode, topics[i].name);
	   }
	}
  }
} 

function createTopics(tarray) 
{
	//alert('Looking for named anchors as topics');
	for (i=0; i<document.anchors.length; i++)
	{
		a = document.anchors[i];
		// Make sure the anchor has a parent whose first character is a H for HTML headers
		if (a.parentNode && a.parentNode.nodeName.toLowerCase().charAt(0)=='h')  {
		   	// Check if class=obe_topic or obe_subtopic and add into the topic table accordingly
           	switch (a.parentNode.className)
		   	{
				case "obe_subtopic":
					if (tarray[tarray.length-1].subtopics==undefined) tarray[tarray.length-1].subtopics = new Array();
					tarray[tarray.length-1].subtopics.push(new Topic(a.name, a.innerHTML));
					break;
				case "obe_topic":
				case "obe_section":
					tarray.push(new Topic(a.name, a.innerHTML));
			   		break;			
				default:
					// ignore all others
					break;
			} // endswitch
		} // endif a.parentNode
	} // endfor
}

function createTOC(tarray)
{
	document.getElementById("obe_toc").innerHTML=printTopics(tarray, true);
}
 
function preLoadImages()
{
  eyeglass.src = "library/show_image.gif";
  eyeglass.alt = "Show Screenshot for Step";
  images = document.images;
  // Only load images with class="imgborder_off"
  for (i = 0; i < images.length; i++) {
    if (images[i].className.substr(0, 9) == "imgborder") {
        img = new Image();
        img.id = images[i].id;
		img.alt = images[i].id;
        img.src = "images/" + images[i].id + ".gif";
        obeImages[img.id] = img;
    }
  }
}
 
function hideImage(obj)
{
    obj.src = eyeglass.src;
	obj.alt = eyeglass.alt;
    obj.className = "imgborder_off";
}

function showImage(obj)
{
    newImg = obeImages[obj.id];
    obj.src = newImg.src;
//	obj.alt = newImg.alt;
	obj.alt = "Screenshot for Step";
    obj.className = "imgborder_on";
}

function toggleImage(obj)
{
    if (obj.src == eyeglass.src) {
		showImage(obj);
    }
    else {
		hideImage(obj);
    }
}

function toggleAllImages(obj) 
{
  if (obj.src.indexOf("showall") > -1) {
     obj.src = obj.src.replace("showall", "hideall");
	 obj.alt = "Click icon to hide all screenshots";
	 showAllImages();
  }
  else {
     obj.src = obj.src.replace("hideall","showall");
	 obj.alt ="Click icon to show all screenshots";
	 hideAllImages();
  }
  document.getElementById("togglealltext").innerHTML=obj.alt;
}
 
function hideAllImages()
{
  imgs = document.images;
  for (i=0; i < imgs.length; i++)
  {
    if (imgs[i].className.substr(0, 9) == "imgborder") {
	  hideImage(imgs[i]);
	}
  }
}

function showAllImages()
{
  imgs = document.images;
  for (i=0; i < imgs.length; i++)
  {
    if (imgs[i].className.substr(0, 9) == "imgborder") {
	  showImage(imgs[i]);
	}
  }
}

function toggleVisible(id)
{
  var e = document.getElementById(id);
  if (isNav) {
    e.style.display = (e.style.display == "table-row") ? "none" : "table-row";
  }
  else {
    e.style.display = (e.style.display == "block") ? "none" : "block";
  }
}

function toggleSubtopics(obj) {
  //alert("toggleSubtopics: obj.alt=" + obj.alt);
  if (obj.alt == "expand") {
    obj.alt= "collapse";
  }
  else {
    obj.alt = "expand";
  } 
  obj.src = "library/" + obj.alt + ".gif"
}